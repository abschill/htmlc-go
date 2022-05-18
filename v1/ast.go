package upv1

/**
 * Partials
**/
const ASTPartialReggie = "<!--@partial=([a-z | 0-9 | _ | -]+)-->"

/**
 * Keys
**/
const ASTKeyReggie = "<!--@render=([a-z | 0-9 | _ | -]+)-->"

/**
 * Iterators
**/
const ASTLoopOpenReggie = "<!--@loop=(.*?)"
const ASTLoopCloseReggie = "\\)-->"
const ASTLoopReggie = "<!--@loop=(.*?)[\n*\t*]*?(<\\w+>.*)\n?\\)-->"
