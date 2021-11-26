#include <stdio.h>
#include <stdlib.h>

int mystrlen(char *);
char *mystradd(char *,char *);
int mystrfind(char *,char *);

int main(int argc, char *argv[]) {
  char *busqueda1 = (char *)malloc(sizeof(char *));
  busqueda1=argv[2];
  char *busqueda2 = (char *)malloc(sizeof(char *));
  busqueda2=argv[3];
  if(mystrfind(argv[1],"-find")==0 && argc ==4){
     printf("[%s] string was found at [%d] position\n", busqueda2,mystrfind(busqueda1,busqueda2));
  }
  else if(mystrfind(argv[1],"-add")==0 && argc ==4){
    printf("Initial Lenght      : %d\n", mystrlen(busqueda1));
    mystradd(busqueda1,busqueda2);
    printf("New String          :%s\n",busqueda1);
    printf("New length          :%d\n",mystrlen(busqueda1));
  } 
  else{
      printf("error en parametros\n");
      return -1;
  }
  return 0;
}
