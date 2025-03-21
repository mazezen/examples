package com.mazezen.springbootmybatis.controller;

import com.mazezen.springbootmybatis.pojo.User;
import com.mazezen.springbootmybatis.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserController {

    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @RequestMapping("/findById")
    public User findById(Integer id) {
       return  userService.findById(id);
    }

}
