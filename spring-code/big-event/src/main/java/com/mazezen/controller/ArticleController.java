package com.mazezen.controller;

import com.mazezen.pojo.Article;
import com.mazezen.response.Result;
import com.mazezen.service.impl.ArticleServiceImpl;
import com.mazezen.utils.ThreadLocalUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/article")
public class ArticleController {

    @Autowired
    private ArticleServiceImpl articleService;

    @PostMapping("/add")
    public Result<String> addArticle(@RequestBody @Validated Article article) {
        Map<String, Object> map = ThreadLocalUtil.get();
        Integer uid = (Integer) map.get("id");

        Article articleByTitle = articleService.findByTile(uid, article.getTitle());
        if (articleByTitle != null) {
            return Result.error("文章名称重复");
        }
        article.setCreateTime(LocalDateTime.now());
        article.setUpdateTime(LocalDateTime.now());
        article.setCreateUser(uid);
        articleService.add(article);
        return Result.success("添加成功");
    }

    @GetMapping("/list")
    public Result<List<Article>> list() {
        Map<String, Object> map = ThreadLocalUtil.get();
        Integer uid = (Integer) map.get("id");

        List<Article> list = articleService.list(uid);
        return Result.success(list);
    }

    @DeleteMapping("/delete")
    public Result<String> deleteArticle(@RequestParam("id") Integer id) {
        Map<String, Object> map = ThreadLocalUtil.get();
        Integer uid = (Integer) map.get("id");
        Article article = articleService.findById(id);
        if (article == null) {
            return Result.error("文章不存在");
        }

        articleService.delete(uid, id);
        return Result.success();
    }

}
