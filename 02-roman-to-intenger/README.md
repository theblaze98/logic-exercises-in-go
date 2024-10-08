# 13. Romano a Entero

### Nivel
**Fácil** :3rd_place_medal:

## Pista
Los números romanos están representados por siete símbolos diferentes: I, V, X, L, C, D y M.

| Símbolo | Valor |
|---------|-------|
| I       | 1     |
| V       | 5     |
| X       | 10    |
| L       | 50    |
| C       | 100   |
| D       | 500   |
| M       | 1000  |

Por ejemplo, 2 se escribe como II en números romanos, simplemente dos unos sumados. 12 se escribe como XII, que es simplemente X + II. El número 27 se escribe como XXVII, que es XX + V + II.

Los números romanos generalmente se escriben de mayor a menor de izquierda a derecha. Sin embargo, el número cuatro no se escribe IIII. En su lugar, el número cuatro se escribe como IV. Debido a que el uno está antes del cinco, lo restamos, haciendo cuatro. El mismo principio se aplica al número nueve, que se escribe como IX. Hay seis instancias donde se usa la resta:

- I puede colocarse antes de V (5) y X (10) para hacer 4 y 9.
- X puede colocarse antes de L (50) y C (100) para hacer 40 y 90.
- C puede colocarse antes de D (500) y M (1000) para hacer 400 y 900.

Dado un número romano, conviértelo a un entero.

### Ejemplo 1:

**Entrada:** `s = "III"`
**Salida:** `3`
**Explicación:** III = 3.

### Ejemplo 2:

**Entrada:** `s = "LVIII"`
**Salida:** `58`
**Explicación:** L = 50, V = 5, III = 3.

### Ejemplo 3:

**Entrada:** `s = "MCMXCIV"`
**Salida:** `1994`
**Explicación:** M = 1000, CM = 900, XC = 90 e IV = 4.

### Restricciones:

- 1 <= s.length <= 15
- s contiene solo los caracteres ('I', 'V', 'X', 'L', 'C', 'D', 'M').
- Se garantiza que s es un número romano válido en el rango [1, 3999].
