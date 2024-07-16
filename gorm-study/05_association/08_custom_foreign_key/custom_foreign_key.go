package _8_custom_foreign_key

//当未使用标准的字段进行关联时，需要对关联属性进行设置。
//
//推荐尽量采用标准的模型定义。
//
//典型的需要自定义的情况：
//
//- 复合主键
//- 数据库结构已定
//- 多个关联，例如，Essay和Author关联了多次，有第一作者，校订作者，通讯作者等。
//### 外键字段
//
//使用gorm标签：foreignKey来自定义外键字段，要求与关联字段类型一致。
//
//### 引用字段
//
//使用gorm标签：references来自定义引用字段，要求与外键字段类型一致。
//
//### 约束操作
//
//使用gorm标签：constraint来自定义约束操作：
//
//- OnUpdate
//  - CASCADE，级联更新
//  - SET NULL，外键设置NULL
//  - RESTRICT，限制更新
//- OnDelete
//  - CASCADE，级联删除
//  - SET NULL，外键设为NULL
//  - RESTRICT，限制删除，默认
