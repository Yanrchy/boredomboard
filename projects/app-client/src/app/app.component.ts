import { Component } from '@angular/core';
import { TranscriptComponent } from './transcript/transcript.component';
import { LobbyComponent } from './lobby/lobby.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [ TranscriptComponent, LobbyComponent ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'app-client';
}
