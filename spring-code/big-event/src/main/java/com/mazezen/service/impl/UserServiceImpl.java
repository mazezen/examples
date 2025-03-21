package com.mazezen.service.impl;

import com.mazezen.mapper.UserMapper;
import com.mazezen.pojo.User;
import com.mazezen.service.IUserService;
import com.mazezen.utils.Md5Util;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
public class UserServiceImpl implements IUserService {

    @Autowired
    private UserMapper userMapper;

    @Override
    public User findByUsername(String username) {
        return userMapper.findByUsername(username);
    }

    @Override
    public void add(String username, String password) {
        String md5String = Md5Util.getMD5String(password);
        userMapper.add(username, md5String);
    }

    @Override
    public User findById(Integer uid) {
        return userMapper.findById(uid);
    }

    @Override
    public void updateInfo(User user) {
        user.setUpdateTime(LocalDateTime.now());
        userMapper.updateInfo(user);
    }

    @Override
    public void updateAvatar(Integer uid, String avatar) {
        userMapper.updateAvatar(uid, avatar);
    }

    @Override
    public void updatePwd(Integer uid, String md5String) {
        userMapper.updatePwd(uid, md5String);
    }
}
