import { Component } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { FormGroup, FormControl } from '@angular/forms';
import { Router } from '@angular/router';

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
  m_RouterService:  Router;

  m_UserForm = new FormGroup({
    username  : new FormControl(''),
    userColor : new FormControl('')
  });


  constructor(profileService: UserProfileService, router: Router) {

    this.m_ProfileService = profileService;
    this.m_RouterService = router;

  }

  Connect() {

    this.m_ProfileService.Register(
      this.m_UserForm.value.username as string,
      this.m_UserForm.value.userColor as string
    );

    this.m_RouterService.navigateByUrl("/comm-channel");

  }

}
