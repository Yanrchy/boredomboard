import { Component } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { FormGroup, FormControl } from '@angular/forms';

import { UserProfileService } from '../user-profile.service';

@Component({
  
  selector: 'app-connection-prompt',
  standalone: true,
  imports: [ ReactiveFormsModule ],
  templateUrl: './connection-prompt.component.html',
  styleUrl: './connection-prompt.component.scss',

  providers: [ UserProfileService ]

})
export class ConnectionPromptComponent {

  m_ProfileService: UserProfileService;

  m_UserForm = new FormGroup({
    username  : new FormControl(''),
    userColor : new FormControl('')
  });

  constructor(profileService: UserProfileService) {

    this.m_ProfileService = profileService;

  }

  Connect() {

    this.m_ProfileService.Register(
      this.m_UserForm.value.username as string,
      this.m_UserForm.value.userColor as string
    );

  }

}
