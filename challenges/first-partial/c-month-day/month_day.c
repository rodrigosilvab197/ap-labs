#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int day[] = {
    31, 28, 31,30, 31, 30,31, 31, 30,31, 30, 31
};
int month[] = {
    1, 2, 3,4, 5, 6,7, 8, 9,10,11,12
};
char *month_name[] = {
    "Jan", "Feb", "Mar","Apr", "May", "Jun","Jul", "Aug", "Sep","Oct", "Nov", "Dec"
};
/* month_day function's prototype*/
void month_day(int year, int yearday, int *pmonth, int *pday){
    int i = 0;
    int j = 0;
    pmonth = &month[0];
    pday = &day[0];
     if (year%4 == 0 && year >0){
        if(yearday>0 && yearday<367){
            pday[1] = 29;
            while(j<12){
                if((yearday-pday[i])>0){
                 yearday=yearday-pday[i];
                    i++;
                }
                j++;
            }
            printf("%s %02d, %d\n", month_name[i], yearday, year);
        }
        else{
                printf(" por favor debes pasar un dia valido! Tu año es bisisesto maximo 366");
        }

        }
        else if (year%4 != 0 && year >0){
            if(yearday>0 && yearday <366){
                while(j <12){
                    if((yearday-pday[i])>0){
                        yearday=yearday-pday[i];
                        i++;
                    }
                    j++;
                 }
                 
                printf("%s %02d, %d\n", month_name[i], yearday, year);
            }    
            
            else{
                printf("por favor debes pasar un dia valido! maximo 365");
            }
        
        }
        else{
            printf("por favor debes ingresar un año que sea un año mayor a 0");
        }
    
   

}

int main(int argc, char *argv[]) {

    if(3 != argc){
        printf(" por favor debes pasar dos argumentos , el año y el dia!");
        return -1;
    }
    else{
            int year = atoi(argv[1]);
            int day = atoi(argv[2]);
            int pmonth, pday;
            month_day(year, day, &pmonth, &pday);
       
    }
    return -1;
}
