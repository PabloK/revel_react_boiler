# mkdir $GOPATH\build\rrb-<buildno>
# revel test rrb
# cp test-results $GOPATH\build\rrb-<buildno>
# pid = revel run rrb&
# npm test jasmine1
# pkill -9 pid 
# revel build rrb $GOPATH\build\rrb-<buildno>
