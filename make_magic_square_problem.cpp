// Magic Square
// This is a program for making new magic square problem automatically.
// example>
//    b b d h
//  + c c c d
//  ----------  --> Your kids have to solve this problem by replacing them into natural numbers.
//  i f i c h

// 1. Randomize number list  { 3, 1, 2, 4, 6, 9, 0, 8, 5, 7 }
//    number_list[0] = 3, number_list[1] = 1, ...
// 2. Prepare alphabet list
//    { 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j' }
// 3. Map alphabets to numbers
//    number_list[0] = 3 = alpha_list[0] = 'a'  ==> (3, 'a')
//    number_list[1] = 1 = alpha_list[1] = 'b'  ==> (1, 'b')
// 4. Get two random nubmers (1st and 2nd)
// 5. Calculate the sum of the two numbers. this is 3rd numbers to be used.
//    So 1st + 2nd = 3rd
// 6. Change this numbers into characters. Save them into some buffers.
//    These will be used for showing answers to stdout later.
// 7. Replace 3 numbers into mapping alphabets. This is the problem to solve.

#include <cstdio>
#include <iostream>
#include <cstdlib>
#include <ctime>

#define NUM_COUNT  10

using namespace std;

void swap( int num_list[], int aidx, int bidx )
{
    int  tmp;

    tmp = num_list[aidx];
    num_list[aidx] = num_list[bidx];
    num_list[bidx] = tmp;
}

void show_num_list( int num_list[] )
{
    for( int i = 0; i < NUM_COUNT; i++ )
    {
        cout << i << " ";
    }
    cout << endl;

    for( int i = 0; i < NUM_COUNT; i++ )
    {
        cout << num_list[i] << " ";
    }
    cout << endl;
}

void show_alpha_list( char alpha_list[] )
{
    for( int i = 0; i < NUM_COUNT; i++ )
    {
        cout << alpha_list[i] << " ";
    }
    cout << endl;
}

void show_magic_number( char alpha_list[], int num_list[], char * str )
{
    char  tmp[4];
    int   idx;

    for( int i = 0; i < strlen(str); i++ )
    {
        tmp[0] = *(char *)(str + i);
        tmp[1] = '\0';

        idx = atoi((char *)tmp);
        printf("%c ", alpha_list[num_list[idx]]);
    }
    printf("\n");
}

int main(int argc, char *argv[])
{
    int  num_list[NUM_COUNT];
    char alpha_list[NUM_COUNT];
    int  i, rnum;
    int  addee1, addee2, result_num;
    char a1[8];
    char a2[8];
    char r[8];

    for( int i = 0; i < NUM_COUNT; i++ )
    {
        num_list[i] = i;
    }

    // randomize num list
    srandom(time(NULL));

    for( i = 0; i < NUM_COUNT - 1; i++ )
    {
        rnum = rand() % ( NUM_COUNT - i );
        swap( num_list, rnum, 9-i );
    }
    show_num_list(num_list);

    // prepare alphabet list for converting numbers to alphabets
    // num_list[i] == alpha_list[i]
    alpha_list[0] = 'a';
    for( i = 1; i < NUM_COUNT; i++ )
    {
        alpha_list[i] = alpha_list[i-1] + 1;
    }
    show_alpha_list(alpha_list);

    // prepare problem numbers
    addee1 = rand() % 10000;
    addee2 = rand() % 10000;
    result_num = addee1 + addee2;

    sprintf( a1, "%d", addee1 );
    sprintf( a2, "%d", addee2 );
    sprintf( r, "%d", result_num );

    printf("addee1(%s) addee2(%s) result(%s)\n", a1, a2, r);

    show_magic_number( alpha_list, num_list, a1 );
    show_magic_number( alpha_list, num_list, a2 );
    show_magic_number( alpha_list, num_list, r );

    return 0;
}
