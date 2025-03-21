package com.mazezen.service;

import com.mazezen.pojo.Category;
import com.mazezen.response.Result;

import java.util.List;

public interface ICategoryService {
    Category findByName(String categoryName);

    Category findByAlias(String categoryAlias);

    void add(Category category);

    Category findById(Integer id, Integer integer);

    void update(Category category);

    List<Category> listCategory(Integer uid);

    void deleteById(Integer uid, Integer id);
}
