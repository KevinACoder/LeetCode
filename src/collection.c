#include <stdlib.h>
#include <stdio.h>

#define bool int
#define false 0
#define true 1
#define nil  NULL
#define size_t unsigned int
#define ARR_ITEM_NUM(arr) (sizeof(arr)/sizeof(arr[0]))

//=====vecotr
typedef struct vector {
    void    **item;
    size_t  len;
    size_t  capcity;
} vector;
vector* vector_init();
vector* vector_append(vector *vec, void *item);
static vector* update_capcity(vector *vec);
void vector_print(vector *vec);
vector* vector_preappend(vector *vec, void *item);

#define VEC_INIT(vec) vector *(vec) = vector_init()
#define VEC_APPEND(vec, item) (vec) = vector_append((vec), (void*)(item))
#define VEC_PREAPPEND(vec, item) (vec) = vector_preappend((vec), (void*)(item))

vector* vector_init()
{
    vector *vec = malloc(sizeof(vector));
    vec->len = 0;
    vec->capcity = 10;
    vec->item = malloc(sizeof(void*)*vec->capcity);
    return vec;
}

static vector* update_capcity(vector *vec)
{
    if(vec->len == vec->capcity){
        vec->capcity *= 2;
        vec->item = realloc(vec->item, vec->capcity*sizeof(void*));
    }
    return vec;
}

vector* vector_append(vector *vec, void *item)
{
    vec = update_capcity(vec);
    vec->item[vec->len++] = item;
    return vec;
}

void vector_print(vector *vec)
{
    for(int i = 0; i < vec->len; ++i){
        printf("%d ", (int)vec->item[i]);
    }
    printf("\n");
}

vector* vector_preappend(vector *vec, void *item)
{
     vec = update_capcity(vec);
     for(int i = vec->len; i >= 0; --i){
         vec->item[i] = vec->item[i-1];
     }
     vec->item[0] = item;
     vec->len++;
     return vec;
}

bool vector_test()
{
    //vector *vec = vector_init();
    VEC_INIT(vec);
    for(int i = 0; i <= 50; ++i){
        //vec = vector_append(vec, (void*)i);
        VEC_APPEND(vec, i);
    }
    for(int i = -1; i >= -10; --i){
        VEC_PREAPPEND(vec, i);
    }
    vector_print(vec);
    return true;
}
//=====vecotr


typedef bool (*test_api_func)();
int main(char *argv[], int argc)
{
    test_api_func funcs[] = {vector_test};
    size_t len = ARR_ITEM_NUM(funcs);
    for(size_t i = 0; i < len; ++i){
        if(!funcs[i]()){
            printf("fail\n");
            return 1;
        }
    }
    printf("pass\n");
    return 0;
}