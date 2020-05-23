/* Less Efficient Solution */

bool backspaceCompare(char * S, char * T) {
    int i = 0, k = 0;
    char *cpyS = calloc(1, strlen(S)+1);
    char *cpyT = calloc(1, strlen(T)+1);
    if (!cpyS || !cpyT) {
        perror("Failed");
        return false;
    }
    
    for (i = 0; i < strlen(S) && k < strlen(S); i++) {
        if (S[i] != '#') {
            cpyS[k++] = S[i];
        } else if (k > 0) {
            k -= 1;
        }
    }
    if (k < strlen(S))
        cpyS[k] = '\0';
    
    k = 0;
    for (i = 0; i < strlen(T) && k < strlen(T); i++) {
        if (T[i] != '#') {
            cpyT[k++] = T[i];
        } else if (k > 0) {
            k -= 1;
        }
    }
    if (k < strlen(T))
        cpyT[k] = '\0';
    bool comparison = !strcmp(cpyS, cpyT);
    free(cpyS);
    free(cpyT);
    return comparison;
}

/* More Efficient solution */

bool backspaceCompare(char * S, char * T) {
    int i = 0, k = 0;
    for (i = 0; i < strlen(S); i++) {
        if (i > 0 && S[i] == '#') {
            for (k = i+1; k <= strlen(S); k++) {
                S[k-2] = S[k];
                S[k] = ' ';
            }
            i -= 2;
        }
    }
    for (i = 0; i < strlen(T); i++) {
        if (i > 0 && T[i] == '#') {
            for (k = i+1; k <= strlen(T); k++) {
                T[k-2] = T[k];
                T[k] = ' ';
            }
            i -= 2;
        }
    }
    for (; *S != '\0' && (*S == '#' || *S == ' '); *S++);
    for (; *T != '\0' && (*T == '#' || *T == ' '); *T++);
    for (i = strlen(S)-1; i >= 0 && (S[i] == '#' || S[i] == ' '); i--);
    S[i+1] = '\0';
    for (i = strlen(T)-1; i >= 0 && (T[i] == '#' || T[i] == ' '); i--);
    T[i+1] = '\0';
    return !strcmp(S, T);
}

