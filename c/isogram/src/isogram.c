#include "isogram.h"
#include <ctype.h>
#include <string.h>

bool is_isogram(const char phrase[]) {
  if (phrase == NULL)
    return false;
  unsigned long i;
  unsigned long j;
  unsigned int offset = 0;
  char checkphrase[strlen(phrase)];
  for (i = 0; i < strlen(phrase); i++) {
    if (phrase[i] == '-' || phrase[i] == ' ')
      ++offset;
    else
      checkphrase[i] = tolower(phrase[i + offset]);
  }
  for (i = 0; i < strlen(checkphrase); i++) {
    for (j = i + 1; j < strlen(checkphrase); j++) {
      if (checkphrase[i] == checkphrase[j])
        return false;
    }
  }
  return true;
}
