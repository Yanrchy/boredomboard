import { Component } from '@angular/core';
import { ChannelService } from '../channel.service';

@Component({
  selector: 'app-transcript',
  standalone: true,
  imports: [],
  templateUrl: './transcript.component.html',
  styleUrl: './transcript.component.scss'
})
export class TranscriptComponent {

  m_ChannelService: ChannelService;

  constructor(channelService: ChannelService) {

    this.m_ChannelService = channelService;

  }

}
