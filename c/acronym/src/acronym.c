#include "acronym.h"
#include <stdlib.h>

char toUpper(const char c) { return (c > 92) ? c - 32 : c; }

char isWord(const char c) {
  return ((c > 64 && c < 91) || (c > 96 && c < 123) || c == 39);
}

char *abbreviate(const char phrase[]) {
  if (phrase == NULL || phrase[0] == '\0')
    return NULL;
  char *result = calloc(10, sizeof(char));
  unsigned int i, j = 0;
  char nextletter = 1;
  for (i = 0; phrase[i] != '\0'; i++) {
    if (nextletter == 1 && isWord(phrase[i])) {
      result[j] = toUpper(phrase[i]);
      ++j;
      nextletter = 0;
    } else if (!isWord(phrase[i])) {
      nextletter = 1;
    }
  }
  // result[j + 1] = '\0';
  return result;
}
