package com.mazezen.springbootmybatis.service.impl;

import com.mazezen.springbootmybatis.mapper.UserMapper;
import com.mazezen.springbootmybatis.pojo.User;
import com.mazezen.springbootmybatis.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserServiceImpl implements UserService {

    private final UserMapper userMapper;

    public UserServiceImpl(UserMapper userMapper) {
        this.userMapper = userMapper;
    }

    @Override
    public User findById(Integer id) {
        return userMapper.findById(id);
    }
}
