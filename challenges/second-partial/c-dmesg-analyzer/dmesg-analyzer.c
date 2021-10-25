#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define REPORT_FILE "report.txt"
#define HASHSIZE 100000

void analizeLog(char *logFile, char *report);

// Structures
struct lista{
    struct lista *next;
    char *problema;
    char *nombre;

};

static struct lista *hashtable[HASHSIZE];
struct lista* element;
unsigned hash(char *s){
    unsigned hashvalue;
    // For loop
    for (hashvalue = 0; *s != '\0'; s++){
        hashvalue = *s + 31 * hashvalue;
    }
    return hashvalue%HASHSIZE;
}
struct lista *buscar(char *s){
    struct lista *tmp;
    // For loop
    for (tmp = hashtable[hash(s)]; tmp != NULL; tmp = tmp->next){
        if(strcmp(s, tmp->nombre) == 0){
            return tmp;
        }
    }
    return NULL;
}



struct lista *hacerLista(char *nombre, char *problema){
    struct lista *tmp;
    unsigned hashvalue;
    // Conditional part
    if ((tmp = buscar(nombre)) == NULL){
        tmp = (struct lista *) malloc(sizeof(*tmp));
        if(tmp == NULL || (tmp->nombre = strdup(nombre))==NULL){
            return NULL;
        }
        hashvalue = hash(nombre);
        tmp->next = hashtable[hashvalue];
        hashtable[hashvalue] = tmp;
    }
    else{
        char *newLength;
        newLength = tmp->problema;
        tmp->problema = malloc(strlen(tmp->problema) + strlen(problema) + 50000);
        strcpy(tmp->problema, newLength);
        strcat(tmp->problema, "    ");
        strcat(tmp->problema, problema);
        return NULL;
    }

    if((tmp->problema = strdup(problema)) == NULL){
        return NULL;
    }
    return tmp;
}
void crear(char *x1,char *x2,char *x3, char *x4, char *copiar_linea){
           char * newLine = malloc(strlen(x2) + strlen(x4) +4);
            strcpy(newLine,x2);
            strcat(newLine,"]");
            strcat(newLine,x4);
            if(strcmp(x1, x3) == 0){
                hacerLista("General", copiar_linea);
            }else{
                hacerLista(x1, newLine);
            }
        }

void analizeLog(char *logFile, char *report) {
    printf("Generating Report from: [%s] log file\n", logFile);
    size_t bufer = 0;
    char * copiarLinea;
    ssize_t contador;
    char * linea = NULL;
    FILE * archivo = fopen(logFile, "r");
    if (archivo == NULL){
        printf("File not found\n");
        exit(EXIT_FAILURE);
    }
    contador = getline(&linea, &bufer, archivo);
    while (contador >= 0){
        bufer++;
        copiarLinea = strdup(linea);
        char *x1 = strtok(linea, "]");
        char *x2 = strdup(x1);
        x1 = strtok(NULL, ":");
        char *x4 = strtok(NULL, "");
        if (x4 == NULL)
        {
        x4 = x1;
        }
        char *x3 = strtok(copiarLinea, "]");
        x3 = strtok(NULL, "");
        if(x1 != NULL){
            crear(x1,x2,x3,x4,copiarLinea);
        } 
        contador = getline(&linea, &bufer, archivo);   
          }
    fclose(archivo);
    archivo = fopen(REPORT_FILE, "w");
    if(archivo == NULL){
        printf("Error creating a file");
        exit(EXIT_FAILURE);
    }
    for (int i = 0; i < HASHSIZE; i++){
        if(hashtable[i]!=NULL){
         fprintf(archivo, "%s\n", hashtable[i]->nombre);
         fprintf(archivo, "    %s\n", hashtable[i]->problema);
        }
    }
    printf("Report is generated at: [%s]\n", report);
}

int main(int argc, char **argv) {
    if (argc < 2) {
	printf("Faltan argumentos\n");
	return 1;
    }
    analizeLog(argv[1], REPORT_FILE);
    return 0;
}
