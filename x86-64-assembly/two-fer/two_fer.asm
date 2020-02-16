default rel

section .rodata
msgpre:    db "One for ", 0
msgdef:     db "you", 0
msgsuf:     db ", one for me.", 0

section .text
global two_fer
two_fer:
    push rdi
    lea rdi, [msgpre]
    call cpz
    pop rdi
    cmp rdi, 0
    jne .name
    lea rdi, [msgdef]
.name:
    call cpz
    lea rdi, [msgsuf]
    jmp cpz

__cpz_prelude:
    inc rsi
    inc rdi
cpz:
    mov al, [rdi]
    mov [rsi], al
    cmp al, 0
    jne __cpz_prelude
    ret
