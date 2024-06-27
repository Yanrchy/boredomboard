import { Component } from '@angular/core';

import { TranscriptComponent } from '../transcript/transcript.component';
import { LobbyComponent } from '../lobby/lobby.component';
import { InputFieldComponent } from '../input-field/input-field.component';

import { ChannelService } from '../channel.service';

@Component({
  selector: 'app-comm-channel',
  standalone: true,
  imports: [TranscriptComponent, LobbyComponent, InputFieldComponent],
  templateUrl: './comm-channel.component.html',
  styleUrl: './comm-channel.component.scss',
  providers: [ChannelService]
})
export class CommChannelComponent {

  m_ChannelService: ChannelService;

  constructor(channelService: ChannelService) {

    this.m_ChannelService = channelService;

  }

}
