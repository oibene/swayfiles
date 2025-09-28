import Image from "next/image"

import person from "../../../public/logos/person.svg"

export default function ModalAccount({open}){
     {/* CLICK OUTSIDE BUTTON AREA */}

    const ref = useRef<HTMLDivElement>(null);
    useEffect(() => {
        const handleClickOut = (event: MouseEvent) => {
            if ( open &&
                ref.current && !ref.current.contains(event.target)
            ){
                setOpenAcc(false);
            }
        };

        document.addEventListener('mousedown', handleClickOut);

        return () => {
            document.removeEventListener('mousedown', handleClickOut);
        };

    }, [open]);

    return(
        <div className={open ? "absolute bg-dark-gray w-70 h-52" : "hidden"}>
            <div className="font-noto font-bold text-light-gray text-sm">

                <div className="m-3">
                    <p>CONTA</p>

                    <div className="flex mt-2 content-center ">
                        <div className="flex h-14 w-14 justify-center rounded-full bg-light-gray">
                            <Image src={person} alt=""></Image>
                        </div>
                        <div className="ml-3 text-xs content-center">
                            <p>Nome Conta</p>
                            <p className="font-normal">email@email.com</p>
                        </div>
                    </div>
                </div>
               
                <hr className="w-full h-0.5 rounded-sm"/>

                <button className="m-3 cursor-pointer">
                    <p>ALTERAR CONTA</p>
                </button>

                <hr className="w-full h-0.5 rounded-sm"/>

                <button className="m-3 cursor-pointer">
                    <p>LOGOUT</p>
                </button>
            </div>

        </div>
    )
}