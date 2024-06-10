import { Component } from '@angular/core';
import { TranscriptComponent } from './transcript/transcript.component';
import { LobbyComponent } from './lobby/lobby.component';
import { InputFieldComponent } from './input-field/input-field.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [ TranscriptComponent, LobbyComponent, InputFieldComponent ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'app-client';
}
