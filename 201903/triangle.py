#-*- coding:utf-8 -*-

import os, sys

def hollow_triangle(level):
    print("空心三角形")
    #边长/层数
    int_level = int(level)
    for floor in range(1, int_level+1):
        #输出空格
        for i in range(int_level-floor):
            print(" ",end="")
        #输出星号
        all_star = floor*2-1
        for j in range(all_star):
            if j == 0 or j == all_star-1:
                print("*",end="")
            elif floor == int_level:
                if j%2 == 0:
                    print("*",end="")
                else:
                    print(" ",end="")
            else:
                print(" ",end="")
        print("")

if __name__ == "__main__":
    print("Start")
    level = sys.argv[1]
    hollow_triangle(level)
    print("End")
