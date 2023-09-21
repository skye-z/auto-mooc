(function () {
    'use strict';
    const GT = 500;
    let runWork = () => {
        let extraTime = 0;
        try {
            let done = false;
            let video = document.querySelector('#player_html5_api');
            if (video) {
                video.playbackRate = 2;
                video.muted = true;
                if (!video.ended) video.play();
                else video.pause();
                if (video.ended) done = true;
                let quizLayer = document.querySelector('#quizLayer');
                if (quizLayer && quizLayer.style.display != 'none') {
                    if (done) {
                        setTimeout(() => {
                            document.querySelectorAll('.layui-layer-shade').forEach(e => e.style.display = 'none');
                        }, GT << 1);
                    };
                    let source = JSON.parse(document.querySelector('div[uooc-video]').getAttribute('source'));
                    let quizList = source.quiz;
                    let quizIndex = 0;
                    let quizQuestion = document.querySelector('.smallTest-view .ti-q-c').innerHTML;
                    for (let i = 0; i < quizList.length; i++) {
                        if (quizList[i].question == quizQuestion) {
                            quizIndex = i;
                            break;
                        };
                    };
                    let quizAnswer = eval(quizList[quizIndex].answer);
                    let quizOptions = quizLayer.querySelector('div.ti-alist');
                    for (let ans of quizAnswer) {
                        let labelIndex = ans.charCodeAt() - 'A'.charCodeAt();
                        quizOptions.children[labelIndex].click();
                    };
                    quizLayer.querySelector('button').click();
                    extraTime = 1000;
                };
                if (!done) {
                    if (video.paused) video.play();
                    else document.querySelectorAll('.layui-layer-shade, #quizLayer').forEach(e => e.style.display = 'none');
                };
            };
            if (!done) setTimeout(runWork, GT + extraTime);
            else if (video) {
                let current_video = document.querySelector('.basic.active');
                let next_part = current_video.parentNode;
                let next_video = current_video;
                let isVideo = node => Boolean(node.querySelector('span.icon-video'));
                let canBack = () => {
                    return Boolean(next_part.parentNode.parentNode.tagName === 'LI');
                };
                let toNextVideo = () => {
                    next_video = next_video.nextElementSibling;
                    while (next_video && !isVideo(next_video)) {
                        next_video = next_video.nextElementSibling;
                    };
                };
                let isExistsVideo = () => {
                    let _video = next_part.firstElementChild;
                    while (_video && !isVideo(_video)) {
                        _video = _video.nextElementSibling;
                    };
                    return Boolean(_video && isVideo(_video));
                };
                let isExistsNextVideo = () => {
                    let _video = current_video.nextElementSibling;
                    while (_video && !isVideo(_video)) {
                        _video = _video.nextElementSibling;
                    };
                    return Boolean(_video && isVideo(_video));
                };
                let isExistsNextListAfterFile = () => {
                    let part = next_part.nextElementSibling;
                    return Boolean(part && part.childElementCount > 0);
                };
                let toNextListAfterFile = () => {
                    next_part = next_part.nextElementSibling;
                };
                let toOuterList = () => {
                    next_part = next_part.parentNode.parentNode;
                };
                let toOuterItem = () => {
                    next_part = next_part.parentNode;
                };
                let isExistsNextListAfterList = () => {
                    return Boolean(next_part.nextElementSibling);
                };
                let toNextListAfterList = () => {
                    next_part = next_part.nextElementSibling;
                };
                let expandList = () => {
                    next_part.firstElementChild.click();
                };
                let toExpandListFirstElement = () => {
                    next_part = next_part.firstElementChild.nextElementSibling;
                    if (next_part.classList.contains('unfoldInfo')) {
                        next_part = next_part.nextElementSibling;
                    };
                };
                let isList = () => {
                    return Boolean(next_part.tagName === 'UL');
                };
                let toInnerList = () => {
                    next_part = next_part.firstElementChild;
                };
                let toFirstVideo = () => {
                    next_video = next_part.firstElementChild;
                    while (next_video && !isVideo(next_video)) {
                        next_video = next_video.nextElementSibling;
                    };
                };
                let mode = {
                    FIRST_VIDEO: '1',
                    NEXT_VIDEO: '2',
                    LAST_LIST: '3',
                    NEXT_LIST: '4',
                    INNER_LIST: '5',
                    OUTER_LIST: '6',
                    OUTER_ITEM: '7',
                }
                let search = (_mode) => {
                    switch (_mode) {
                        case mode.FIRST_VIDEO:
                            if (isExistsVideo()) {
                                toFirstVideo();
                                next_video.click();
                                setTimeout(runWork, GT);
                            } else if (isExistsNextListAfterFile()) search(mode.LAST_LIST);
                            break;
                        case mode.NEXT_VIDEO:
                            if (isExistsNextVideo()) {
                                toNextVideo();
                                next_video.click();
                                setTimeout(runWork, GT);
                            } else if (isExistsNextListAfterFile()) search(mode.LAST_LIST);
                            else search(mode.OUTER_ITEM);
                            break;
                        case mode.LAST_LIST:
                            toNextListAfterFile();
                            toInnerList();
                            search(mode.INNER_LIST);
                            break;
                        case mode.NEXT_LIST:
                            toNextListAfterList();
                            search(mode.INNER_LIST);
                            break;
                        case mode.INNER_LIST:
                            expandList();
                            (function waitForExpand() {
                                if (next_part.firstElementChild.nextElementSibling) {
                                    toExpandListFirstElement();
                                    if (isList()) {
                                        toInnerList();
                                        search(mode.INNER_LIST);
                                    } else search(mode.FIRST_VIDEO);
                                } else setTimeout(waitForExpand, GT);
                            })();
                            break;
                        case mode.OUTER_LIST:
                            toOuterList();
                            if (isExistsNextListAfterList()) search(mode.NEXT_LIST);
                            else if (canBack()) search(mode.OUTER_LIST);
                            break;
                        case mode.OUTER_ITEM:
                            toOuterItem();
                            if (isExistsNextListAfterList()) {
                                toNextListAfterList();
                                search(mode.INNER_LIST);
                            } else if (canBack()) search(mode.OUTER_LIST);
                            break;
                        default:
                            break;
                    };
                };
                try {
                    search(mode.NEXT_VIDEO);
                } catch (err) {
                    console.error(err);
                };
            };
        } catch (err) {
            console.error(err);
        };
    };
    runWork();
    console.info('Auto Mooc Init');
})();

function selectClass() {
    let lv1 = document.querySelector('.basic.active');
    let lv1Part = lv1.parentNode;
    let lv1Sub = lv1Part.querySelector('ul');
    let lv1Res = lv1Part.querySelector('.resourcelist');
    if (lv1Res != null) {
        let res = lv1Res.firstElementChild
        res.click()
    } else if (lv1Sub != null) {
        let lv2Part = lv1Sub.firstElementChild
        let lv2 = lv2Part.querySelector('.basic');
        lv2.click()
        setTimeout(() => {
            let lv2Sub = lv2Part.querySelector('ul');
            let lv2Res = lv2Part.querySelector('.resourcelist');
            if (lv2Res != null) {
                let res = lv2Res.firstElementChild
                res.click()
            } else if (lv2Sub != null) {
                let lv3Part = lv2Sub.firstElementChild
                let lv3 = lv3Part.querySelector('.basic');
                lv3.click()
            }
        }, 500)
    }
}