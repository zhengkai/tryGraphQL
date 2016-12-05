var lQuery = {}
var iQuery = 7;
var iQueryCountdown = iQuery;

function syntaxHighlight(json) {
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
        var cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            } else {
                cls = 'string';
            }
        } else if (/true|false/.test(match)) {
            cls = 'boolean';
        } else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}

function loadQuery(i) {
	$.get({
		url: 'query/' + (i < 10 ? '0' + i : i) + '.txt',
	}).done(function(data) {
		lQuery[i] = data.replace(/\t/g, '    ');
	});
	init();
}

function init() {
	if (iQueryCountdown > 0) {
		iQueryCountdown--;
		return;
	}

	if (iQueryCountdown < 0) {
		return;
	}
	iQueryCountdown = -1;

	console.log('done')
	for (var i = 1; i <= iQuery; i++) {
		var o = $('<button class="btn btn-outline-primary"></button>').data('i', i).text(i).click(function() {
			var i = $(this).data('i');
			$('#queryBox').val(lQuery[i]);
		}).appendTo('#queryBar');

		if (i == 6) {
			o.trigger('click');
			$('#submitButton').trigger('click');
		}
	}
}

function sendQuery() {
	var s = $('#queryBox').val();

	$.post({
		url: 'https://soulogic.com/graphql/api',
		data: s,
	}).done(function(data) {
		var json = JSON.stringify(data, undefined, 4);
		$('#resultBox').html(syntaxHighlight(json));
	});
}

$(document).ready(function () {
	init();
});

(function () {
	for (var i = 1; i <= iQuery; i++) {
		loadQuery(i);
	}
})()
