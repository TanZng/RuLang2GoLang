fact=1.0;
sum = 1.0;
i = 1.0;
n = 11.0;
x = 1.0;
while (i < n) {
    fact = fact*i;
    sum = sum + ((x^i)/fact);
    i = i + 1.0;
}

imprime("La suma de la serie de Taylor es :");
imprime(sum);
