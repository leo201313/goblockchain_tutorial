rd /s /q tmp
md tmp\blocks
main.exe createblockchain -address LeoCao
main.exe blockchaininfo
main.exe balance -address LeoCao
main.exe send -from LeoCao -to Krad -amount 100
main.exe balance -address Krad
main.exe mine
main.exe blockchaininfo
main.exe balance -address LeoCao
main.exe balance -address Krad
main.exe send -from LeoCao -to Exia -amount 100
main.exe send -from Krad -to Exia -amount 30
main.exe mine
main.exe blockchaininfo
main.exe balance -address LeoCao
main.exe balance -address Krad
main.exe balance -address Exia