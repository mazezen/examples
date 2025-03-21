package com.mazezen.service.impl;

import com.mazezen.mapper.ArticleMapper;
import com.mazezen.pojo.Article;
import com.mazezen.service.IArticleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ArticleServiceImpl implements IArticleService {
    @Autowired
    private ArticleMapper articleMapper;

    @Override
    public Article findByTile(Integer uid, String title) {
        return articleMapper.findByTile(uid, title);
    }

    @Override
    public void add(Article article) {
        articleMapper.add(article);
    }

    @Override
    public List<Article> list(Integer uid) {
        return articleMapper.list(uid);
    }

    @Override
    public void delete(Integer uid, Integer id) {
        articleMapper.delete(uid, id);
    }

    @Override
    public Article findById(Integer id) {
        return articleMapper.findById(id);
    }
}
