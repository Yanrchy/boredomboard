import { Route } from "@angular/router";
import { CommChannelComponent } from "./comm-channel/comm-channel.component";
import { ConnectionPromptComponent } from "./connection-prompt/connection-prompt.component";

export const routes : Route[] = [
    { path: "comm-channel", component: CommChannelComponent },
    { path: "prompt",       component: ConnectionPromptComponent },

    { path: "", redirectTo: "/prompt", pathMatch: "full" }
];
