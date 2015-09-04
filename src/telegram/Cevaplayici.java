/*
 * Copyright (C) 2015 erkanmdr
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
 */
package telegram;

/**
 *
 * @author erkanmdr
 */
import io.github.nixtabyte.telegram.jtelebot.client.RequestHandler;
import io.github.nixtabyte.telegram.jtelebot.exception.JsonParsingException;
import io.github.nixtabyte.telegram.jtelebot.exception.TelegramServerException;
import io.github.nixtabyte.telegram.jtelebot.request.TelegramRequest;
import io.github.nixtabyte.telegram.jtelebot.request.factory.TelegramRequestFactory;
import io.github.nixtabyte.telegram.jtelebot.response.json.Message;
import io.github.nixtabyte.telegram.jtelebot.server.impl.AbstractCommand;
import java.util.logging.Level;
import java.util.logging.Logger;
import static telegram.Cevirici.andeyoConverter;
import static telegram.Cevirici.latinConverter;
import static telegram.Cevirici.latin_alfabesi;

public class Cevaplayici extends AbstractCommand {

    public Cevaplayici(Message message, RequestHandler requestHandler) {
        super(message, requestHandler);
    }

    /**
     *
     */
    @Override
    public void execute() {
        String s;
        try {

            if (message.getChat().getId() == 81329453 || message.getChat().getId() == 119941223) {

                if (Cevirici.latinmi(latin_alfabesi, message.getText())) {
                    s = andeyoConverter(message.getText());

                } else {
                    s = latinConverter(message.getText());

                }
            } else {
                s = "Meraklı taze, burada ne işin var? Hadi yaylanda boyunu görelim....";
            }

            //if message.getChat().getId()
            TelegramRequest telegramRequest = TelegramRequestFactory.createSendMessageRequest(message.getChat().getId(), s, true, null, null);
            requestHandler.sendRequest(telegramRequest);

        } catch (JsonParsingException | TelegramServerException ex) {
            Logger.getLogger(Cevaplayici.class.getName()).log(Level.SEVERE, null, ex);
        }

    }

}
