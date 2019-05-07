function maskify( a) {
    if (a.length <= 4){
        return a
    }
    for ( i = 0; i < a.length - 4; i++) {
        a[i]= '#'
    }
    return a
}