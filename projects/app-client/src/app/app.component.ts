import { Component } from '@angular/core';
import { TranscriptComponent } from './transcript/transcript.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [ TranscriptComponent ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'app-client';
}
