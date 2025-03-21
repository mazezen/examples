/**
 * TCP客户端通信基本流程
 */

#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <iostream>
#include <string.h>

#define SERVER_ADDRESS "127.0.0.1"
#define SERVER_PORT 3000
#define SEND_DATA "helloworld"

int main(int argc, char *argv[])
{
    // 1. create socket
    int clientfd = socket(AF_INET, SOCK_STREAM, 0);
    if (clientfd == -1)
    {
        std::cout << "create client socket error:" << std::endl;
        return -1;
    }

    // 2. connect server
    struct sockaddr_in serverAddr;
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = inet_addr(SERVER_ADDRESS);
    serverAddr.sin_port = htons(SERVER_PORT);
    if (connect(clientfd, (struct sockaddr *)&serverAddr, sizeof(serverAddr)) == -1)
    {
        std::cout << "connect socket error." << std::endl;
        return -1;
    }

    // 3. send data to server
    int ret = send(clientfd, SEND_DATA, strlen(SEND_DATA), 0);
    if (ret != strlen(SEND_DATA))
    {
        std::cout << "send data error." << std::endl;
        return -1;
    }
    std::cout << "send data successfully, data: " << SEND_DATA << std::endl;

    // 4. recv data from client
    char recvBuf[32] = {0};
    ret = recv(clientfd, recvBuf, 32, 0);
    if (ret > 0)
    {
        std::cout << "recv data successfully, data: " << recvBuf << std::endl;
    }
    else
    {
        std::cout << "recv data error, data: " << recvBuf << std::endl;
    }

    // 5. 关闭socket
    // close(clientfd);
    // 这里仅仅是为了让客户端程序不退出
    while (true)
    {
        sleep(3);
    }
    return 0;
}