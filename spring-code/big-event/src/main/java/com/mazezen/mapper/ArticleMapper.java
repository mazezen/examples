package com.mazezen.mapper;

import com.mazezen.pojo.Article;
import org.apache.ibatis.annotations.Delete;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;

import java.util.List;

@Mapper
public interface ArticleMapper {

    @Select("select * from article where create_user=#{uid} and title=#{title}")
    Article findByTile(Integer uid, String title);

    @Insert("insert into article(category_id,create_user,state,title,content,category_alias,cover_img,create_time,update_time) " +
            "values(#{categoryId},#{createUser},#{state},#{title},#{content},#{categoryAlias},#{coverImg},#{createTime},#{updateTime})")
    void add(Article article);

    @Select("select * from article where create_user=#{uid}")
    List<Article> list(Integer uid);

    @Delete("delete from article where create_user=#{uid} and id=#{id}")
    void delete(Integer uid, Integer id);

    @Select("select * from article where id=#{id}")
    Article findById(Integer id);
}
