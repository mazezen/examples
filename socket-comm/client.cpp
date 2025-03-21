/**
 * TCP客户端通信基本流程
 */

#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <iostream>
#include <string.h>
#include <stdio.h>
#include <fcntl.h>
#include <errno.h>

#define SERVER_ADDRESS "127.0.0.1"
#define SERVER_PORT 3000
#define SEND_DATA "helloworld"

// g++ -g -o client client.cpp -std=c++17
int main(int argc, char *argv[])
{
    // 1. create socket
    int clientfd = socket(AF_INET, SOCK_STREAM, 0);
    if (clientfd == -1)
    {
        std::cout << "create client socket error:" << std::endl;
        return -1;
    }

    int oldSocketFlag = fcntl(clientfd, F_GETFL, 0);
    int newSocketFlag = oldSocketFlag | O_NONBLOCK;
    if (fcntl(clientfd, F_SETFL, newSocketFlag) == -1)
    {
        close(clientfd);
        std::cout << "set socket to nonbloack error." << std::endl;
        return -1;
    }

    // 2. connect server
    struct sockaddr_in serverAddr;
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = inet_addr(SERVER_ADDRESS);
    serverAddr.sin_port = htons(SERVER_PORT);
    for (;;)
    {
        int ret = connect(clientfd, (struct sockaddr *)&serverAddr, sizeof(serverAddr));
        if (ret == 0)
        {
            std::cout << "connect to server successfully." << std::endl;
            close(clientfd);
            return 0;
        }
        else if (ret == -1)
        {
            if (errno == EINTR)
            {
                std::cout << "connecting interruptted bu singal, try again." << std::endl;
                continue;
            }
            else if (errno == EINPROGRESS)
            {
                break;
            }
            else
            {
                close(clientfd);
                return -1;
            }
        }
    }

    fd_set writeset;
    FD_ZERO(&writeset);
    FD_SET(clientfd, &writeset);
    // 可以利用tv_sec和tv_usec做更小精度的超时控制
    struct timeval tv;
    tv.tv_sec = 3;
    tv.tv_usec = 0;
    if (select(clientfd + 1, NULL, &writeset, NULL, &tv) != 1)
    {
        std::cout << "[select] connect to server error." << std::endl;
        close(clientfd);
        return -1;
    }

    int err;
    socklen_t len = static_cast<socklen_t>(sizeof err);
    if (::getsockopt(clientfd, SOL_SOCKET, SO_ERROR, &err, &len) < 0)
    {
        close(clientfd);
        return -1;
    }

    if (err == 0)
        std::cout << "connect to server successfully." << std::endl;
    else
        std::cout << "connect to server error." << std::endl;

    // 5. 关闭socket
    close(clientfd);

    return 0;
}