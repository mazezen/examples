package com.mazezen.mapper;

import com.mazezen.pojo.User;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface UserMapper {

    @Select("select * from user where username=#{username}")
    User findByUsername(String username);

    @Insert("insert into user(username,password,create_time,update_time) values " +
            "(#{username},#{password},now(),now())")
    void add(String username, String password);

    @Select("select * from user where id=#{id}")
    User findById(Integer uid);

    @Update("update user set username=#{username}, nickname=#{nickname}, email=#{email}, update_time=#{updateTime} where id=#{id}")
    void updateInfo(User user);

    @Update("update user set user_pic=#{avatar}, update_time=now() where id=#{uid}")
    void updateAvatar(Integer uid, String avatar);

    @Update("update user set password=#{md5String} where id=#{uid}")
    void updatePwd(Integer uid, String md5String);
}
