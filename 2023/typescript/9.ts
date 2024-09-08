const input = `7 10 23 62 147 303 562 967 1585 2545 4127 6933 12171 22103 40826 75939 142603 273520 542191 1113540 2347099
1 6 26 67 139 256 437 720 1221 2314 5086 12343 30619 73880 169929 370916 769849 1525598 2899594 5308259 9396171
21 27 33 40 47 50 37 -24 -207 -663 -1671 -3708 -7541 -14344 -25843 -44492 -73683 -117993 -183471 -277968 -411513
9 14 19 24 29 34 39 44 49 54 59 64 69 74 79 84 89 94 99 104 109
22 47 97 192 368 686 1245 2203 3807 6432 10636 17262 27671 44287 71794 119567 206268 368021 672227 1239924 2280674
16 42 94 185 328 536 822 1199 1680 2278 3006 3877 4904 6100 7478 9051 10832 12834 15070 17553 20296
2 6 22 72 190 419 808 1409 2274 3452 4986 6910 9246 12001 15164 18703 22562 26658 30878 35076 39070
1 7 26 63 118 186 257 316 343 313 196 -43 -444 -1052 -1917 -3094 -4643 -6629 -9122 -12197 -15934
28 41 54 67 80 93 106 119 132 145 158 171 184 197 210 223 236 249 262 275 288
21 36 57 78 95 113 153 259 505 1002 1905 3420 5811 9407 14609 21897 31837 45088 62409 84666 112839
9 8 21 66 162 319 527 753 965 1229 1977 4623 12811 34737 87308 203745 448540 949561 1970920 4080259 8525175
13 29 56 113 236 495 1038 2190 4644 9798 20332 41201 81368 156845 295987 548539 1000723 1799735 3192476 5585253 9633656
1 10 22 37 55 76 100 127 157 190 226 265 307 352 400 451 505 562 622 685 751
6 17 35 60 92 131 177 230 290 357 431 512 600 695 797 906 1022 1145 1275 1412 1556
8 23 57 125 242 423 683 1037 1500 2087 2813 3693 4742 5975 7407 9053 10928 13047 15425 18077 21018
15 34 70 134 237 390 604 890 1259 1722 2290 2974 3785 4734 5832 7090 8519 10130 11934 13942 16165
-6 5 41 125 290 576 1034 1760 2995 5351 10282 21061 44848 97132 211244 458332 988036 2108385 4436971 9176709 18601746
19 34 72 153 305 564 974 1587 2463 3670 5284 7389 10077 13448 17610 22679 28779 36042 44608 54625 66249
10 1 -11 -12 27 157 469 1141 2544 5449 11389 23242 46113 88605 164581 295531 513670 865905 1418821 2264848 3529783
22 34 46 58 70 82 94 106 118 130 142 154 166 178 190 202 214 226 238 250 262
3 14 34 63 110 219 517 1288 3083 6902 14541 29293 57338 110394 210669 400247 759734 1447500 2785930 5454466 10920055
15 25 58 127 245 425 680 1023 1467 2025 2710 3535 4513 5657 6980 8495 10215 12153 14322 16735 19405
0 14 37 77 156 318 658 1386 2940 6162 12551 24607 46280 83538 145068 243124 394536 621894 954921 1432049 2102212
-3 -12 -24 -36 -41 -29 16 144 540 1794 5513 15560 40353 96834 216906 457303 913955 1741868 3181266 5590120 9482071
13 7 -9 -37 -76 -122 -168 -204 -217 -191 -107 57 326 728 1294 2058 3057 4331 5923 7879 10248
10 23 59 137 287 563 1084 2133 4366 9225 19735 42025 88189 181566 366296 724326 1405266 2678207 5020667 9269445 16873011
15 16 18 37 107 294 711 1528 2979 5394 9346 16138 29129 56915 120304 266635 597838 1323846 2862891 6029551 12389270
5 12 19 26 33 40 47 54 61 68 75 82 89 96 103 110 117 124 131 138 145
23 35 53 74 96 131 222 471 1102 2603 6025 13593 29963 64852 138576 293595 618033 1292159 2677229 5479680 11044906
19 39 63 86 97 71 -36 -275 -656 -1045 -995 493 5310 16841 41340 91336 193721 407543 860307 1817312 3806879
3 16 43 90 167 288 471 738 1115 1632 2323 3226 4383 5840 7647 9858 12531 15728 19515 23962 29143
13 37 74 124 187 263 352 454 569 697 838 992 1159 1339 1532 1738 1957 2189 2434 2692 2963
10 35 76 137 220 325 450 591 742 895 1040 1165 1256 1297 1270 1155 930 571 52 -655 -1580
23 40 57 76 107 178 352 750 1586 3244 6482 12968 26609 56657 124591 278661 623412 1377796 2987436 6340494 13186306
-2 3 15 50 140 329 670 1230 2120 3587 6247 11624 23321 49424 107175 231597 490670 1010903 2018787 3905714 7326574
19 38 65 94 116 127 148 266 713 2023 5355 13152 30430 67164 142466 291538 576735 1104492 2050357 3694930 6474136
23 49 85 125 156 161 130 79 77 281 979 2641 5978 12009 22136 38227 62707 98657 149921 221221 318280
-1 12 47 123 265 504 877 1427 2203 3260 4659 6467 8757 11608 15105 19339 24407 30412 37463 45675 55169
11 29 55 93 155 270 494 921 1695 3023 5189 8569 13647 21032 31476 45893 65379 91233 124979 168389 223507
4 26 75 168 333 630 1185 2250 4326 8425 16610 33063 66139 132253 263146 519269 1013960 1956092 3723349 6984746 12900057
16 24 41 75 135 231 391 710 1445 3164 6942 14574 28757 53211 92817 154129 247178 390472 622683 1026909 1776847
11 25 61 142 302 582 1035 1756 2968 5232 9924 20254 43319 94039 202436 426812 877431 1759273 3450277 6647180 12642226
11 20 53 131 289 586 1128 2116 3931 7267 13321 24046 42469 73071 122220 198641 313899 482862 724101 1060173 1517721
22 39 67 108 163 240 385 749 1702 4010 9121 19704 40833 82741 167090 340500 702046 1456089 3008835 6137288 12273887
14 39 79 134 204 289 389 504 634 779 939 1114 1304 1509 1729 1964 2214 2479 2759 3054 3364
5 20 57 133 271 509 921 1649 2948 5255 9308 16362 28576 49678 86054 148452 254543 432638 726923 1204643 1965741
18 32 46 60 74 88 102 116 130 144 158 172 186 200 214 228 242 256 270 284 298
5 16 43 95 180 309 523 958 1963 4286 9343 19585 38978 73611 132447 228232 378577 607228 945539 1434163 2124976
11 17 26 44 89 197 446 1013 2294 5154 11451 25124 54390 116010 243221 499863 1004541 1970447 3768832 7027180 12777023
11 15 36 81 162 305 574 1120 2269 4676 9594 19338 38064 73032 136580 249103 443407 770893 1310120 2178399 3547182
23 31 45 91 221 531 1193 2513 5040 9774 18552 34723 64243 117309 210579 369856 632801 1050725 1687723 2614271 3890813
3 4 3 -6 -32 -87 -186 -347 -591 -942 -1427 -2076 -2922 -4001 -5352 -7017 -9041 -11472 -14361 -17762 -21732
4 17 43 85 150 249 397 613 920 1345 1919 2677 3658 4905 6465 8389 10732 13553 16915 20885 25534
5 15 25 42 96 261 688 1654 3638 7449 14452 26966 48943 87079 152557 263678 449699 756267 1252915 2043170 3277914
21 38 58 89 158 322 679 1379 2635 4734 8048 13045 20300 30506 44485 63199 87761 119446 159702 210161 272650
1 9 32 73 142 278 592 1353 3159 7261 16140 34475 70684 139270 264260 484087 858333 1476825 2471656 4032789 6427994
2 -7 -25 -62 -129 -234 -378 -551 -728 -865 -895 -724 -227 756 2424 5019 8830 14197 21515 31238 43883
10 38 90 179 321 531 815 1158 1508 1756 1712 1077 -589 -3903 -9699 -19072 -33426 -54526 -84554 -126169 -182571
16 44 99 204 391 706 1220 2051 3407 5666 9516 16186 27808 47960 82451 140421 235842 389520 631713 1005496 1571021
10 28 49 88 175 352 677 1254 2314 4381 8585 17253 35057 71303 144563 292069 588586 1183618 2372905 4731846 9356939
21 34 50 76 120 201 380 820 1888 4337 9680 21035 45042 96013 204385 432964 906635 1864632 3749962 7359690 14095184
16 21 34 76 181 398 797 1479 2586 4299 6800 10157 14062 17293 16648 4857 -33464 -129507 -344518 -795362 -1698921
18 22 23 24 33 71 190 501 1212 2676 5449 10358 18579 31725 51944 82027 125526 186882 271563 386212 538805
2 7 13 30 83 224 553 1253 2653 5344 10386 19659 36428 66211 118060 206388 353500 593013 974379 1568756 2476505
21 33 54 107 229 474 916 1652 2805 4527 7002 10449 15125 21328 29400 39730 52757 68973 88926 113223 142533
14 33 74 149 273 480 846 1522 2802 5295 10359 21136 44886 98015 216542 477386 1041025 2236169 4726501 9836181 20178546
22 41 71 112 164 227 301 386 482 589 707 836 976 1127 1289 1462 1646 1841 2047 2264 2492
22 45 82 139 222 337 490 687 934 1237 1602 2035 2542 3129 3802 4567 5430 6397 7474 8667 9982
16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36
12 26 40 54 68 82 96 110 124 138 152 166 180 194 208 222 236 250 264 278 292
18 32 47 70 128 292 724 1772 4152 9269 19753 40350 79473 152114 285715 532521 995850 1885250 3629311 7099026 14035066
6 10 17 20 19 39 166 617 1861 4805 11052 23227 45352 83232 144791 240270 382168 584772 863083 1230902 1697793
26 54 98 163 255 386 585 915 1496 2534 4356 7451 12517 20514 32723 50811 76902 113654 164342 232947 324251
16 27 37 48 61 74 84 105 231 802 2787 8609 23842 60552 143568 321665 686482 1402873 2755081 5213274 9524043
-3 -4 -4 -3 -1 2 6 11 17 24 32 41 51 62 74 87 101 116 132 149 167
10 24 51 112 244 506 983 1795 3139 5425 9621 18024 35875 74618 158291 335721 703168 1445282 2907415 5720568 11013222
11 22 50 110 224 422 758 1353 2476 4673 8953 17039 31691 57107 99407 167204 272265 430264 661628 992476 1455650
-2 0 2 -2 -25 -90 -217 -376 -376 368 3295 11555 31511 75216 164268 335493 648906 1198328 2124864 3633140 6009719
26 39 60 102 198 427 967 2197 4886 10529 21912 43998 85207 159098 286355 496918 832369 1349962 2133335 3323373 5199807
16 33 65 112 174 251 343 450 572 709 861 1028 1210 1407 1619 1846 2088 2345 2617 2904 3206
0 0 1 -3 -19 -50 -76 -5 427 1833 5460 13618 30294 61997 118885 216230 376282 630598 1022907 1612587 2478835
2 -1 -6 -9 0 43 172 514 1366 3371 7814 17085 35364 69591 130792 235840 409738 688519 1122866 1782563 2761896
-9 -11 -16 -29 -55 -99 -166 -261 -389 -555 -764 -1021 -1331 -1699 -2130 -2629 -3201 -3851 -4584 -5405 -6319
5 19 50 112 224 416 747 1342 2465 4657 8986 17493 33996 65564 125235 236975 444517 826647 1522794 2774518 4990768
17 20 36 80 171 341 653 1231 2305 4274 7790 13866 24011 40395 66047 105089 163009 246976 366200 532340 759963
5 15 42 109 244 486 916 1743 3503 7471 16457 36283 78464 165031 337215 671186 1306767 2499885 4717698 8808367 16296866
15 22 30 37 52 112 319 914 2415 5859 13201 27933 55990 107005 195958 345232 587039 966108 1542432 2393749 3617280
12 20 29 39 50 62 75 89 104 120 137 155 174 194 215 237 260 284 309 335 362
-4 -8 -6 20 108 330 818 1819 3801 7637 14902 28340 52620 95649 171032 302954 536278 958149 1745431 3270162 6332958
11 25 53 101 175 281 425 613 851 1145 1501 1925 2423 3001 3665 4421 5275 6233 7301 8485 9791
4 9 27 81 219 532 1175 2391 4538 8119 13815 22521 35385 53850 79699 115103 162672 225509 307267 412209 545271
5 -1 -9 -2 59 252 727 1767 3894 8048 15888 30303 56300 102619 184881 332199 601807 1112009 2115581 4158820 8414301
0 -4 -8 -3 38 176 529 1308 2887 5964 11940 23760 47637 96330 194983 390967 769714 1479204 2766576 5031295 8900432
7 7 24 84 224 502 1033 2079 4242 8848 18680 39336 81679 166153 330254 641396 1219370 2276906 4196464 7682334 14071065
8 18 56 140 303 602 1124 1989 3350 5390 8316 12350 17717 24630 33272 43775 56196 70490 86480 103824 121979
5 9 13 17 21 25 29 33 37 41 45 49 53 57 61 65 69 73 77 81 85
16 21 42 101 236 507 1011 1929 3653 7082 14243 29506 61853 128976 264481 530249 1037166 1979139 3687780 6718663 11985050
6 13 34 83 191 423 917 1961 4133 8551 17325 34388 67042 128858 245150 463335 872464 1639605 3076316 5757101 10724601
13 20 29 50 98 193 360 629 1035 1618 2423 3500 4904 6695 8938 11703 15065 19104 23905 29558 36158
24 49 85 123 156 182 214 305 605 1499 3941 10210 25481 60838 138668 301776 628062 1253213 2404597 4450413 7969162
19 27 35 43 51 59 67 75 83 91 99 107 115 123 131 139 147 155 163 171 179
3 15 37 60 72 71 84 189 539 1403 3291 7355 16528 38420 92112 223227 536117 1260893 2890627 6458388 14089218
16 23 38 77 172 384 826 1698 3336 6274 11313 19587 32617 52353 81232 122393 180636 266185 407428 688861 1352614
5 21 57 122 219 348 528 853 1605 3464 7882 17726 38345 79279 156905 298407 547565 972983 1679519 2823842 4635223
9 32 69 128 224 379 621 983 1503 2228 3240 4773 7627 14418 32957 84647 223953 580916 1449163 3459586 7908630
1 14 39 75 117 172 288 603 1430 3403 7718 16512 33432 64455 119029 211614 363711 606476 984025 1557545 2410335
18 36 75 141 248 441 833 1653 3308 6475 12264 22541 40588 72428 129386 232827 422550 771070 1407041 2552421 4579720
22 45 73 99 111 92 30 -53 -29 516 2709 9321 26639 67995 160137 354503 747208 1515079 2981981 5742639 10894379
4 7 15 51 166 448 1027 2082 3872 6844 11924 21182 39192 75599 149676 298027 587092 1132765 2130277 3898557 6944602
6 14 22 30 38 46 54 62 70 78 86 94 102 110 118 126 134 142 150 158 166
11 12 8 -2 -17 -34 -48 -52 -37 8 96 242 463 778 1208 1776 2507 3428 4568 5958 7631
27 39 51 63 75 87 99 111 123 135 147 159 171 183 195 207 219 231 243 255 267
-4 5 39 111 236 442 792 1419 2575 4697 8518 15341 27820 52060 102701 214076 462817 1011029 2185874 4624881 9545263
17 27 45 91 203 445 918 1774 3233 5603 9303 14889 23083 34805 51208 73716 104065 144347 197057 265143 352059
15 27 51 93 161 257 363 424 332 -87 -1086 -3021 -6408 -12049 -21280 -36414 -61475 -103345 -173475 -290343 -482877
13 23 28 38 88 251 653 1490 3047 5719 10034 16678 26522 40651 60395 87362 123473 170999 232600 311366 410860
-9 -14 -14 -4 21 66 136 236 371 546 766 1036 1361 1746 2196 2716 3311 3986 4746 5596 6541
13 25 48 92 175 323 570 958 1537 2365 3508 5040 7043 9607 12830 16818 21685 27553 34552 42820 52503
26 49 93 184 373 741 1414 2603 4684 8333 14731 25854 44863 76609 128268 210121 336494 526873 807209 1211428 1783161
6 3 13 58 178 440 947 1859 3450 6252 11392 21324 41308 82205 165454 331487 653334 1257785 2357223 4296134 7617350
17 34 57 82 105 122 129 122 97 50 -23 -126 -263 -438 -655 -918 -1231 -1598 -2023 -2510 -3063
15 19 27 56 152 409 992 2164 4317 8007 13993 23280 37166 57293 85702 124892 177883 248283 340359 459112 610356
12 18 37 92 220 472 918 1660 2857 4768 7829 12831 21432 37647 71817 150148 335622 770418 1756577 3901282 8366744
1 4 21 65 149 286 489 771 1145 1624 2221 2949 3821 4850 6049 7431 9009 10796 12805 15049 17541
-4 5 37 105 225 430 802 1527 2978 5831 11219 20929 37647 65256 109192 176863 278136 425897 636689 931433 1336237
11 34 68 112 168 246 375 636 1261 2880 7038 17136 39973 88117 183546 363719 693227 1289008 2377872 4427638 8438466
10 29 72 153 286 485 764 1137 1618 2221 2960 3849 4902 6133 7556 9185 11034 13117 15448 18041 20910
2 24 63 126 229 413 774 1509 2995 5948 11754 23124 45300 88129 169427 320175 592224 1069336 1882553 3231066 5409951
9 10 22 64 177 451 1079 2453 5324 11057 22033 42306 78763 143373 257898 464183 846811 1582296 3043408 6014896 12129629
15 29 46 62 82 139 326 846 2094 4809 10379 21467 43290 86202 170874 338734 673401 1345847 2711760 5523866 11393778
6 10 10 3 -14 -44 -90 -155 -242 -354 -494 -665 -870 -1112 -1394 -1719 -2090 -2510 -2982 -3509 -4094
6 7 6 3 -2 -9 -18 -29 -42 -57 -74 -93 -114 -137 -162 -189 -218 -249 -282 -317 -354
-3 -8 -7 13 72 197 422 788 1343 2142 3247 4727 6658 9123 12212 16022 20657 26228 32853 40657 49772
17 29 64 134 249 418 650 955 1345 1835 2444 3196 4121 5256 6646 8345 10417 12937 15992 19682 24121
3 4 4 20 96 312 787 1682 3214 5698 9654 16072 27061 47414 88325 176082 372006 817055 1828709 4104839 9139099
28 41 60 101 195 400 820 1638 3171 5952 10844 19217 33322 57270 99624 179775 342342 687287 1428908 3005201 6271333
20 44 80 139 250 468 882 1623 2872 4868 7916 12395 18766 27580 39486 55239 75708 101884 134888 175979 226562
9 17 49 126 293 647 1379 2830 5561 10437 18725 32206 53301 85211 132071 199118 292873 421337 594201 823070 1121701
6 24 49 89 164 307 565 1000 1690 2730 4233 6331 9176 12941 17821 24034 31822 41452 53217 67437 84460
0 5 11 14 5 -30 -110 -259 -506 -885 -1435 -2200 -3229 -4576 -6300 -8465 -11140 -14399 -18321 -22990 -28495
11 11 3 -15 -48 -108 -205 -317 -330 70 1593 5711 15347 36151 78791 162919 323789 622925 1164783 2122045 3773050
17 28 45 89 205 469 992 1927 3494 6067 10433 18459 34625 69263 145007 309157 656917 1374892 2823099 5691543 11302539
19 43 92 181 340 625 1122 1941 3203 5037 7626 11372 17294 27838 48381 89909 173856 341609 675683 1351907 2772941
-1 5 9 11 18 49 138 332 689 1303 2427 4840 10720 25453 61040 142073 315651 667113 1344093 2592169 4806302
16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36
19 34 49 64 79 94 109 124 139 154 169 184 199 214 229 244 259 274 289 304 319
3 0 4 17 35 52 83 230 827 2712 7686 19231 43571 91172 178789 332180 589619 1006352 1660152 2658141 4145059
17 41 78 139 249 455 851 1643 3286 6734 13853 28056 55228 105018 192584 340886 583631 968983 1564160 2461049 3782979
19 37 76 142 237 359 502 656 807 937 1024 1042 961 747 362 -236 -1093 -2259 -3788 -5738 -8171
24 40 64 98 151 260 532 1231 2958 7013 16094 35587 75841 156011 310298 597726 1116980 2028294 3584932 6177456 10394731
13 28 61 123 231 418 760 1427 2776 5524 11077 22178 44239 88161 176332 355167 720540 1467521 2984038 6020866 11989560
15 27 54 101 173 275 412 589 811 1083 1410 1797 2249 2771 3368 4045 4807 5659 6606 7653 8805
-2 1 4 7 10 13 16 19 22 25 28 31 34 37 40 43 46 49 52 55 58
1 9 23 52 107 196 317 449 541 499 171 -670 -2345 -5290 -10079 -17449 -28327 -43859 -65441 -94752 -133789
9 7 5 3 1 -1 -3 -5 -7 -9 -11 -13 -15 -17 -19 -21 -23 -25 -27 -29 -31
10 33 80 170 333 617 1095 1872 3092 4945 7674 11582 17039 24489 34457 47556 64494 86081 113236 146994 188513
17 24 45 99 219 473 999 2054 4077 7766 14169 24789 41703 67695 106403 162480 241769 351492 500453 699255 960531
16 25 38 53 68 81 90 93 88 73 46 5 -52 -127 -222 -339 -480 -647 -842 -1067 -1324
7 12 21 39 67 98 128 203 530 1687 4974 12954 30240 64591 128387 240560 429065 733982 1211347 1937817 3016281
6 30 69 130 230 409 766 1546 3315 7264 15682 32632 64853 122895 222473 386000 644228 1037890 1619195 2452982 3617288
6 23 48 77 102 120 162 359 1087 3284 9134 23517 57025 132108 295317 641093 1356788 2805601 5675320 11238221 21793992
-1 -2 5 42 161 461 1121 2463 5055 9869 18543 33882 60898 108976 196264 358386 665704 1257965 2414882 4700675 9257596
8 24 57 116 211 345 511 706 979 1539 2962 6564 15084 34010 74265 157614 327021 665000 1323052 2563184 4806036
5 13 33 78 170 355 730 1488 2990 5873 11196 20606 36465 61807 99879 152848 219011 287509 329099 280958 22756
1 15 38 65 96 159 352 909 2301 5394 11702 23794 45940 85112 152492 265680 451841 752081 1227398 1966615 3096768
13 29 62 127 255 508 994 1882 3417 5935 9878 15809 24427 36582 53290 75748 105349 143697 192622 254195 330743
8 14 20 26 32 38 44 50 56 62 68 74 80 86 92 98 104 110 116 122 128
20 40 78 142 236 360 510 678 852 1016 1150 1230 1228 1112 846 390 -300 -1272 -2578 -4274 -6420
5 16 29 46 74 124 210 348 555 848 1243 1754 2392 3164 4072 5112 6273 7536 8873 10246 11606
8 20 50 112 236 482 967 1907 3671 6839 12251 21029 34549 54335 81842 118090 163106 215126 269504 317270 343274
24 45 85 151 252 412 707 1345 2815 6148 13383 28450 58926 119581 239462 475712 939774 1845663 3596429 6934929 13200124
7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27
8 9 22 67 184 446 972 1940 3600 6287 10434 16585 25408 37708 54440 76722 105848 143301 190766 250143 323560
8 13 27 70 186 468 1093 2367 4780 9071 16303 27948 45982 72990 112281 168013 245328 350497 491075 676066 916098
-7 -2 11 32 61 98 143 196 257 326 403 488 581 682 791 908 1033 1166 1307 1456 1613
13 42 97 192 341 562 894 1444 2491 4696 9524 20105 42973 91434 191695 393248 785219 1520416 2849029 5165980 9084303
19 18 9 -13 -53 -106 -135 -19 561 2350 6959 17720 41369 91345 194109 400839 808263 1594398 3077749 5812309 10735759
12 8 6 7 23 87 258 621 1282 2358 3962 6183 9061 12557 16518 20637 24408 27076 27582 24503 15987
2 23 54 91 131 181 288 614 1590 4192 10389 23821 50781 101609 192673 349241 609802 1032905 1708608 2778619 4472927
-5 -7 -1 24 83 196 389 695 1155 1819 2747 4010 5691 7886 10705 14273 18731 24237 30967 39116 48899
6 4 4 5 12 44 140 366 833 1756 3634 7743 17363 40592 96360 226516 518848 1148900 2452832 5049769 10040630
7 5 0 -8 -8 39 225 728 1853 4081 8126 15000 26086 43219 68775 105768 157955 229949 327340 456824 626340
16 19 21 30 66 169 425 1035 2457 5659 12540 26614 54117 105797 199790 366179 654086 1142467 1956175 3289334 5438636
-2 -5 -7 -12 -18 8 180 774 2356 6005 13736 29323 59888 118914 231825 446007 848114 1592575 2946027 5352222 9519546
-6 -6 -9 -10 20 148 492 1233 2626 5010 8817 14580 22940 34652 50590 71751 99258 134362 178443 233010 299700
5 19 46 104 229 491 1028 2116 4318 8808 18059 37239 76933 158339 323173 652795 1305838 2593508 5131843 10150185 20112260
5 3 -7 -19 -22 8 140 556 1673 4344 10167 21928 44190 84011 151726 261657 432517 687145 1051045 1548999 2198778
17 32 47 62 77 92 107 122 137 152 167 182 197 212 227 242 257 272 287 302 317
6 14 36 87 184 353 661 1293 2709 5949 13218 29009 62280 130731 269282 546856 1099188 2191612 4336068 8502943 16493572
6 16 35 67 138 324 797 1890 4185 8636 16763 31019 55580 98096 173469 311714 573959 1084957 2099984 4144438 8306493
18 43 89 163 272 423 623 879 1198 1587 2053 2603 3244 3983 4827 5783 6858 8059 9393 10867 12488
-4 -4 -7 -19 -48 -106 -204 -329 -392 -136 1007 4121 11084 24930 50306 94035 165796 278932 451397 706853 1075928
4 14 32 56 93 175 378 841 1786 3548 6636 11861 20598 35350 61084 108562 202508 400558 834408 1793566 3886103
5 11 17 23 29 35 41 47 53 59 65 71 77 83 89 95 101 107 113 119 125
5 9 18 32 51 75 104 138 177 221 270 324 383 447 516 590 669 753 842 936 1035
8 16 38 81 169 366 822 1858 4118 8850 18458 37624 75578 150571 298424 588477 1152968 2241152 4320022 8264571 15721851
0 12 46 123 266 491 801 1204 1793 2956 5839 13280 31581 73702 164776 351345 716632 1406013 2670749 4946327 8999400
2 21 64 150 321 668 1368 2732 5264 9731 17244 29350 48135 76338 117476 175980 257342 368273 516872 712806 967501
21 36 61 108 192 331 546 861 1303 1902 2691 3706 4986 6573 8512 10851 13641 16936 20793 25272 30436
0 5 22 49 77 88 53 -70 -338 -825 -1624 -2849 -4637 -7150 -10577 -15136 -21076 -28679 -38262 -50179 -64823`;

const testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`;

const lines = input
  .split("\n")
  .map((lineString) => lineString.split(" ").map(Number));

function extrapolate(line: number[], sum = line[line.length - 1]) {
  const newLine = line.reduce(
    (result: number[], currentNum, index) => {
      if (index == 0) return result;
      const prevNum = line[index - 1];
      return [...result, currentNum - prevNum];
    },
    []
  );

  // if (!newLine.filter((num) => num != 0).length) return sum;
  if (newLine.every((num) => num == 0)) return sum;

  sum += newLine[newLine.length - 1];
  return extrapolate(newLine, sum);
}

function extrapolate2(line: number[], result = line[0]) {
  let currentLine = line;
  let positive = false;

  function execute() {
    currentLine = currentLine.reduce(
      (result: number[], currentNum, index) => {
        if (index == 0) return result;
        const prevNum = currentLine[index - 1];
        return [...result, currentNum - prevNum];
      },
      []
    );
    if (currentLine.every((num) => num == 0)) return result;

    result += currentLine[0] * (positive ? 1 : -1);
    positive = !positive;
    execute();
  }
  execute();
  return result;
}

const extrapolatedVals = lines.map((l) => extrapolate(l));
const result = extrapolatedVals.reduce((a, b) => a + b);
console.log(result);
const extrapolatedVals2 = lines.map((l) => extrapolate2(l));
const result2 = extrapolatedVals2.reduce((a, b) => a + b);
console.log(result2);
