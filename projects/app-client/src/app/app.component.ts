import { Component } from '@angular/core';

import { CommChannelComponent } from './comm-channel/comm-channel.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [ CommChannelComponent ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'app-client';
}
