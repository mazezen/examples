package com.mazezen.controller;

import com.mazezen.pojo.User;
import com.mazezen.response.Result;
import com.mazezen.service.impl.UserServiceImpl;
import com.mazezen.utils.JwtUtil;
import com.mazezen.utils.Md5Util;
import com.mazezen.utils.ThreadLocalUtil;
import jakarta.validation.constraints.Pattern;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping("/user")
public class UserController {

    @Autowired
    private UserServiceImpl userService;

    @PostMapping("/register")
    public Result register(@Pattern(regexp = "^\\S{5,16}$") String username, @Pattern(regexp = "^\\S{5,16}$") String password) {
        User user = userService.findByUsername(username);
        if (user != null) {
            return Result.error("用户名重复");
        }

        userService.add(username, password);

        return Result.success();
    }

    @PostMapping("/login")
    public Result login(@RequestParam("username") @Pattern(regexp = "^\\S{5,16}$") String username, @RequestParam("password") @Pattern(regexp = "^\\S{5,16}$") String password) {
        User user = userService.findByUsername(username);
        if (user == null) {
            return Result.error("用户名错误");
        }
        if (!Md5Util.checkPassword(password, user.getPassword())) {
            return Result.error("密码错误");
        }

        Map<String, Object> map = new HashMap<>();
        map.put("id", user.getId());
        map.put("username", username);

        String token = JwtUtil.generateToken(map);

        return Result.success("登录成功", token);
    }

    @GetMapping("/userInfo")
    public Result<User> findByUsername() {
        Map<String, Object> map = ThreadLocalUtil.get();
        if (map == null) {
            return Result.error("网络出错,联系技术支持!");
        }
        Integer uid = (Integer) map.get("id");
        User user = userService.findById(uid);
        if (user == null) {
            return Result.error("用户不存在");
        }
        return Result.success(user);
    }

    @PutMapping("/updateInfo")
    public Result updateInfo(@RequestBody @Validated User user) {
        User byId = userService.findById(user.getId());
        if (byId == null) {
            return Result.error("用户不存在");
        }
        userService.updateInfo(user);
        return Result.success();
    }

    @PatchMapping("/updateAvatar")
    public Result updateAvatar(@RequestParam("avatar") String avatar) {
        Map<String, Object> map = ThreadLocalUtil.get();
        if (map == null) {
            return Result.error("网络出错,联系技术支持!");
        }
        Integer uid = (Integer) map.get("id");
        User user = userService.findById(uid);
        if (user == null) {
            return Result.error("用户不存在");
        }
        userService.updateAvatar(uid, avatar);
        return Result.success();
    }

    @PutMapping("/updatePwd")
    public Result updatePwd(
            @RequestParam("oldPwd")  @Pattern(regexp = "^\\S{5,16}$") String oldPwd,
            @RequestParam("password") @Pattern(regexp = "^\\S{5,16}$") String password,
            @RequestParam("repeatPwd") @Pattern(regexp = "^\\S{5,16}$") String repeatPwd
    ) {
        Map<String, Object> map = ThreadLocalUtil.get();
        if (map == null) {
            return Result.error("用户不存在");
        }
        Integer uid = (Integer) map.get("id");
        User user = userService.findById(uid);
        if (user == null) {
            return Result.error("用户不存在");
        }

        if (!Md5Util.checkPassword(oldPwd, user.getPassword())) {
            return Result.error("原密码不正确");
        }

        if (!password.equals(repeatPwd)) {
            return Result.error("两次密码不一致");
        }

        userService.updatePwd(uid, Md5Util.getMD5String(password));

        return Result.success();
    }

}
