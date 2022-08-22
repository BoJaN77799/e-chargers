import { AbstractControl } from '@angular/forms';
import * as moment from 'moment';

export function ReservationDateValidator(control: AbstractControl) {
  if (!control.value) return null;

  var date = moment(control.value);
  // 1 and half hour from now
  var startDateTime = moment((new Date()).getTime() + 5400000);

  const isValid = !(startDateTime > date);

  if (!isValid)
    return { invalidDate: true };

  return null;
}