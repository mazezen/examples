package com.mazezen.service;

import com.mazezen.pojo.Article;

import java.util.List;

public interface IArticleService {
    Article findByTile(Integer uid, String title);

    void add(Article article);

    List<Article> list(Integer uid);

    void delete(Integer uid, Integer id);

    Article findById(Integer id);
}
