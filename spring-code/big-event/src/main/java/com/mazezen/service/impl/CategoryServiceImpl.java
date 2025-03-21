package com.mazezen.service.impl;

import com.mazezen.mapper.CategoryMapper;
import com.mazezen.pojo.Category;
import com.mazezen.response.Result;
import com.mazezen.service.ICategoryService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;

@Service
public class CategoryServiceImpl implements ICategoryService {

    @Autowired
    private CategoryMapper categoryMapper;

    @Override
    public Category findByName(String categoryName) {
        return categoryMapper.findByName(categoryName);
    }

    @Override
    public Category findByAlias(String categoryAlias) {
        return categoryMapper.findByAlias(categoryAlias);
    }

    @Override
    public void add(Category category) {
        categoryMapper.add(category);
    }

    @Override
    public Category findById(Integer uid, Integer id) {
        return categoryMapper.findById(uid, id);
    }

    @Override
    public void update(Category category) {
        category.setUpdateTime(LocalDateTime.now());
        categoryMapper.update(category);
    }

    @Override
    public List<Category> listCategory(Integer uid) {
        return categoryMapper.listCategory(uid);
    }

    @Override
    public void deleteById(Integer uid, Integer id) {
        categoryMapper.deleteById(uid, id);
    }
}
