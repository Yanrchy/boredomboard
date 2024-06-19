import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-connection-prompt',
  standalone: true,
  imports: [ RouterLink ],
  templateUrl: './connection-prompt.component.html',
  styleUrl: './connection-prompt.component.scss'
})
export class ConnectionPromptComponent {

}
