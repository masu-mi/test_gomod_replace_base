module github.com/masu-mi/test_gomod_replace_base

go 1.16

replace github.com/masu-mi/test_gomod_replace_dep_a v0.1.0 => github.com/masu-mi/test_gomod_replace_dep_b v0.2.3

require github.com/masu-mi/test_gomod_replace_dep_a v0.1.0 // indirect
