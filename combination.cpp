// Combination Problem
// This is a program for solving combination problem.
// mCr : Combination of r numbers out of m numbers
// Comb() is the solution function.

#include <cstdio>
#include <cstdlib>

long long facto( long long n )
{
    long long s = 1;

    for( long long i = n; i>0; i-- )
    {
        s *= i;
    }

    return s;
}

long long Answer( long long m, long long r )
{
    long long mfacto = facto(m);
    long long rfacto = facto(r);
    long long mrfacto = facto(m-r);

    return ( mfacto / (rfacto * mrfacto) );
}

long long Comb( long long arr[], long long m, long long r )
{
    long long with, without;

    if( m == r )
    {
        return 1;
    }

    if( r == 1 )
    {
        return m;
    }

    with = Comb( &arr[1], m-1, r-1 );
    without = Comb( &arr[1], m-1, r );

    return with + without;
}

int main(int argc, char *argv[])
{
    long long * arr;
    int   i, m, r;

    if( argc < 2 )
    {
        printf("Usage : %s m r\n", argv[0]);
        exit(0);
    }

    m = atoi(argv[1]);
    r = atoi(argv[2]);

    printf("start -> m : %d, r : %d\n", m, r );

    arr = (long long *)malloc(sizeof(long long) * atoi(argv[1]));
    for( i=0; i<m; i++ )
    {
        arr[i] = i;
    }

    int ret = Comb( arr, m, r );
    printf("ret = %d\n", ret);

    int okayval = Answer(m,r);

    if( ret == okayval )
    {
        printf("goooooooooooood\n");
    }
    else
    {
        printf("baaaaaaaaaaaaad. must be : %d\n", okayval );
    }

    return 0;
}

// 조합 문제를 Recursive 로 접근하여 가장 단순한 처리방법을 생각해봤다.
// 학교 때 배운 조합문제는 다 까먹었었는데, Recursive 방식의 가장 단순한
// 방법을 찾다보니 저절로 공식이 만들어졌다.
// mCr = mCr-1 + m-1Cr-1
