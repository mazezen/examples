package com.mazezen.service;

import com.mazezen.pojo.User;

public interface IUserService {
    User findByUsername(String username);

    void add(String username, String password);

    User findById(Integer uid);

    void updateInfo(User user);

    void updateAvatar(Integer uid, String avatar);

    void updatePwd(Integer uid, String md5String);
}
