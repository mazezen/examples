package com.mazezen.springbootmybatis.service;

import com.mazezen.springbootmybatis.pojo.User;

public interface UserService {

    public User findById(Integer id);

}
