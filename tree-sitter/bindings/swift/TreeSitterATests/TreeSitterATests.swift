import XCTest
import SwiftTreeSitter
import TreeSitterA

final class TreeSitterATests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_a())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading A grammar")
    }
}
