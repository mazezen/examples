package com.mazezen.mapper;

import com.mazezen.pojo.Category;
import com.mazezen.response.Result;
import org.apache.ibatis.annotations.*;

import java.util.List;

@Mapper
public interface CategoryMapper {
    @Select("select * from category where category_name=#{categoryName}")
    Category findByName(String categoryName);

    @Select("select * from category where category_alias=#{categoryAlias}")
    Category findByAlias(String categoryAlias);

    @Insert("insert into category(category_name,category_alias,create_user,create_time,update_time) " +
            "values (#{categoryName},#{categoryAlias},#{createUser},now(),now())")
    void add(Category category);

    @Select("select * from category where create_user=#{uid} and id=#{id}")
    Category findById(Integer uid, Integer id);

    @Update("update category set category_name=#{categoryName}, category_alias=#{categoryAlias}, update_time=#{updateTime} " +
            "where id=#{id} and create_user=#{createUser}")
    void update(Category category);

    @Select("select * from category where create_user=#{uid}")
    List<Category> listCategory(Integer uid);

    @Delete("delete from category where create_user=#{uid} and id=#{id}")
    void deleteById(Integer uid, Integer id);
}
