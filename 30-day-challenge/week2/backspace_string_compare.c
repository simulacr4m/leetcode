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


