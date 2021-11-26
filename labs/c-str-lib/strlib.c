#include <stdlib.h>

int mystrlen(char *str){
    int x = 0;
    int y =0;
    while(str[x]!='\0'){
        x++;
    }
    return x;
}

char *mystradd(char *origin, char *addition){
   char *x = origin + mystrlen(origin);
    while(*addition!='\0') {
        *x++ = *addition++;
    }
    *x = '\0';
    return origin;
}

int mystrfind(char *origin, char *substr){
    int origen = mystrlen(origin);
    int limite = mystrlen(substr);
    int x = 0;
    for(int y = 0;y < origen;y++){
      if(*(origin+y) == *(substr+x)){
        x++;
      }
      else {
        y = y - x + 1;
        x = 0;
      }
      if(x==limite){
          return y - x + 1;
      } 
    }
    return -1;
}
