/**
 * @file A grammar for tree-sitter
 * @author Aaron Kriegman <aaronkplus2@gmail.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "a",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
