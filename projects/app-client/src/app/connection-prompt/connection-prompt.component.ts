import { Component } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { FormGroup, FormControl } from '@angular/forms';

@Component({
  selector: 'app-connection-prompt',
  standalone: true,
  imports: [ ReactiveFormsModule ],
  templateUrl: './connection-prompt.component.html',
  styleUrl: './connection-prompt.component.scss'
})
export class ConnectionPromptComponent {

  userForm = new FormGroup({
    username  : new FormControl(''),
    userColor : new FormControl('')
  });

  Connect() {

    console.warn(this.userForm.value);

  }

}
