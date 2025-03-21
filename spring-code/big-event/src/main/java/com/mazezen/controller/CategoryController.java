package com.mazezen.controller;

import com.mazezen.pojo.Category;
import com.mazezen.response.Result;
import com.mazezen.service.impl.CategoryServiceImpl;
import com.mazezen.utils.ThreadLocalUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/category")
public class CategoryController {

    @Autowired
    private CategoryServiceImpl categoryService;

    @PostMapping("/add")
    public Result addCategory(@RequestBody @Validated Category category) {
        Category byName = categoryService.findByName(category.getCategoryName());
        if (byName != null) {
            return Result.error("分类名重复");
        }
        Category byAlias = categoryService.findByAlias(category.getCategoryAlias());
        if (byAlias != null) {
            return Result.error("分类别名重复");
        }

        Map<String, Object> map = ThreadLocalUtil.get();
        Integer uid = (Integer) map.get("id");
        category.setCreateUser(uid);
        categoryService.add(category);

        return Result.success();
    }

    @GetMapping("/categoryInfo")
    public Result getCategoryInfo(@RequestParam("id") Integer id) {
        Map<String, Object> map = ThreadLocalUtil.get();
        Integer uid = (Integer) map.get("id");
        Category category = categoryService.findById(uid, id);
        return Result.success(category);
    }

    @PutMapping("/update")
    public Result updateCategory(@RequestBody @Validated Category category) {
        Category byName = categoryService.findByName(category.getCategoryName());
        if (byName != null) {
            return Result.error("分类名重复");
        }
        Category byAlias = categoryService.findByAlias(category.getCategoryAlias());
        if (byAlias != null) {
            return Result.error("分类别名重复");
        }
        Map<String, Object> map = ThreadLocalUtil.get();
        Integer uid = (Integer) map.get("id");
        category.setCreateUser(uid);
        categoryService.update(category);
        return Result.success();
    }

    @GetMapping("/list")
    public Result<List<Category>> listCategory() {
        Map<String, Object> map = ThreadLocalUtil.get();
        Integer uid = (Integer) map.get("id");
        List<Category> list = categoryService.listCategory(uid);
        return Result.success(list);
    }

    @DeleteMapping("/delete")
    public Result deleteCategory(@RequestParam("id") Integer id) {
        Map<String, Object> map = ThreadLocalUtil.get();
        Integer uid = (Integer) map.get("id");
        Category category = categoryService.findById(uid, id);
        if (category == null) {
            return Result.error("分类不存在");
        }
        categoryService.deleteById(uid, id);
        return Result.success();
    }

}
