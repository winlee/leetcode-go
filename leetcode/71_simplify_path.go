package leetcode

import "strings"

/**
描述
给你一个字符串 path，表示一个 Unix 风格的文件路径，返回简化后的规范路径。
示例
输入: path = "/home/"
输出: "/home"
输入: path = "/home//foo/"
输出: "/home/foo"
输入: path = "/a/./b/../../c/"
输出: "/c"

Unix 路径规则：

符号
含义
.
当前目录，等于没动
..
上级目录，回退一层
//
多个斜杠等于一个 /
/
根目录

示例 1： /home/ → /home
末尾 / 多余，去掉
示例 2： /home//foo/ → /home/foo
// 变成 /
示例 3： /a/./b/../../c/ → /c
/a/./ → . 表示当前，等于 /a/
/a/b/../../ → .. 回退两次：b 回到 a，a 回到根
最后只剩 /c
关键操作：
按 / 分割路径
遇到 . 或空串 → 跳过
遇到 .. → 弹出栈顶（回退）
其他 → 入栈
最后用 / 连接栈内元素

*/
func simplifyPath(path string) string {
    parts := strings.Split(path, "/")
    stack := []string{}
    
    for _, v := range parts {
        if v == "" || v == "." {
            continue
        }
        if v == ".." {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
            continue
        }
        stack = append(stack, v)
    }
    
    return "/" + strings.Join(stack, "/")
}