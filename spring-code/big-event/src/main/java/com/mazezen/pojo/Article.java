package com.mazezen.pojo;

import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.NotNull;

import java.time.LocalDateTime;

public class Article {

    private Integer id;

    @NotNull
    private Integer categoryId;

    private Integer createUser;

    @NotNull
    private Short state;

    @NotEmpty
    private String title;

    @NotEmpty
    private String content;

    private String categoryAlias;
    private LocalDateTime createTime;
    private LocalDateTime updateTime;

    private String coverImg;

    public Article() {}

    public Article(Integer id, Integer categoryId, Integer createUser, Short state, String title, String content, String categoryAlias, LocalDateTime createTime, LocalDateTime updateTime, String coverImg) {
        this.id = id;
        this.categoryId = categoryId;
        this.createUser = createUser;
        this.state = state;
        this.title = title;
        this.content = content;
        this.categoryAlias = categoryAlias;
        this.createTime = createTime;
        this.updateTime = updateTime;
        this.coverImg = coverImg;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public Integer getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(Integer categoryId) {
        this.categoryId = categoryId;
    }

    public Integer getCreateUser() {
        return createUser;
    }

    public void setCreateUser(Integer createUser) {
        this.createUser = createUser;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public String getCategoryAlias() {
        return categoryAlias;
    }

    public void setCategoryAlias(String categoryAlias) {
        this.categoryAlias = categoryAlias;
    }

    public LocalDateTime getCreateTime() {
        return createTime;
    }

    public void setCreateTime(LocalDateTime createTime) {
        this.createTime = createTime;
    }

    public LocalDateTime getUpdateTime() {
        return updateTime;
    }

    public void setUpdateTime(LocalDateTime updateTime) {
        this.updateTime = updateTime;
    }

    public Short getState() {
        return state;
    }

    public void setState(Short state) {
        this.state = state;
    }

    public String getCoverImg() {
        return coverImg;
    }

    public void setCoverImg(String coverImg) {
        this.coverImg = coverImg;
    }
}
