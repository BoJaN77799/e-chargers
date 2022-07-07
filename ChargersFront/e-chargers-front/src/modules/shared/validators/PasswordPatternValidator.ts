import { AbstractControl } from '@angular/forms';

export function PasswordPatternValidator(control: AbstractControl) {
    if (!control.value) return null;

    const isValid = /^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$/.test(control.value);

    if (!isValid)
        return { invalidPattern: true };

    return null;
}