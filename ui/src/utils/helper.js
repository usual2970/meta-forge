// name2label 将name 转换为 label
// 将名称中的下划线转换为空格，同时每个单个单词的首字母大写
export const name2label= (name) => {
    let label=name.replace(/_/g," ");
    label=label.replace(/( |^)[a-z]/g, L => L.toUpperCase());
    return label;
}