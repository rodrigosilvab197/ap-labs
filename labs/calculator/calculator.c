#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// adds/subtracts/multiplies all values that are in the *values array.
// nValues is the number of values you're reading from the array
// operator will indicate if it's an addition (1), subtraction (2) or
// multiplication (3)
long calc(int operator, int nValues, int *values) {
    int valor = values[0];
    if(operator == 1){
       for(int i = 1; i < nValues ; i++){
           valor+=values[i];
       }
       
       return valor;
    }
    else if (operator==2){
       long valor =0;
       valor = valor +values[0];
       for(int i = 1; i < nValues ; i++){
           valor-=values[i];
       }
      
       return valor;
    }
    else{
       for(int i = 1; i < nValues ; i++){
           valor= valor * values[i];

       }
       
       return valor;
    }
}

int main(int argc, char *argv[]) {
    
    if(4> argc){
        printf(" por favor pasa minimo 3 argumentos!");
        return -1;
    }
    else{
    int values[argc-2];
    for(int i = 2 ; i < argc; i ++){
        values[i-2]= atoi(argv[i]);
    }
    if (strcmp(argv[1],"add")==0){
        return(calc(1,argc-2,values));
        
    }
    else if (strcmp(argv[1],"sub")==0){
        return(calc(2,argc-2,values));
    }
    else if (strcmp(argv[1],"mult")==0){
        return(calc(3,argc-2,values));
    }
    else{
        printf("el argumento de operacion no es valido");
        return -1;
    }
    }
    return -1;
   
}